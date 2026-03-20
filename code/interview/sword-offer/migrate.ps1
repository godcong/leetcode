# Sword Offer Migration Script
# This script migrates old sword-offer structure to new organized structure

$ErrorActionPreference = "Stop"

# Mapping of problem numbers to English names (Offer V1)
$offerV1Mapping = @{
    "10" = "fibonacci"
    "20" = "is-number"
    "22" = "find-kth-node"
    "37" = "serialize-binary-tree"
    "38" = "permutation"
    "42" = "max-sub-array"
    "43" = "count-digit-one"
    "47" = "max-value-of-gifts"
    "52" = "first-common-node"
    "53" = "search-in-sorted-array"
    "58" = "reverse-words"
}

# Mapping for Offer V2 (using LeetCode IDs)
$offerV2Mapping = @{
    "29" = "stream-median"
    "41" = "max-in-sliding-window"
    "69" = "translate-numbers-to-strings"
    "91" = "cut-rope"
    "114" = "n-queens"
    "115" = "longest-increasing-path"
}

Write-Host "=== Starting Sword Offer Migration ===" -ForegroundColor Cyan

# Function to create problem directory and move files
function Move-OfferProblem {
    param(
        [string]$SourcePath,
        [string]$TargetBase,
        [string]$Version,
        [hashtable]$Mapping
    )
    
    $problems = Get-ChildItem -Path $SourcePath -Directory
    
    foreach ($problem in $problems) {
        # Extract problem number from directory name
        $num = $problem.Name -replace '[^\d]', ''
        
        if ([string]::IsNullOrEmpty($num)) {
            Write-Warning "Skipping $($problem.Name): No number found"
            continue
        }
        
        # Get English name from mapping or use default
        $englishName = if ($Mapping.ContainsKey($num)) { 
            $Mapping[$num] 
        } else { 
            "problem-$num"
        }
        
        # Create target directory: {version}/{range}/{num}-{name}
        $rangeStart = [math]::Floor([int]$num / 100) * 100
        $rangeEnd = $rangeStart + 99
        $rangeDir = "{0:D4}-{1:D4}" -f $rangeStart, $rangeEnd
        $targetDir = "{0}/{1}/{2:D3}-{3}" -f $TargetBase, $rangeDir, $num, $englishName
        
        New-Item -ItemType Directory -Path $targetDir -Force | Out-Null
        
        # Move all files from source to target
        $files = Get-ChildItem -Path $problem.FullName -File
        foreach ($file in $files) {
            # Rename files to use simple naming
            $ext = $file.Extension
            $newFileName = "{0}{1}" -f $englishName, $ext
            Move-Item -Path $file.FullName -Destination (Join-Path $targetDir $newFileName) -Force
            Write-Host "  Moved: $($file.Name) -> $newFileName" -ForegroundColor Green
        }
        
        Write-Host "Migrated: $($problem.Name) -> $targetDir" -ForegroundColor Yellow
    }
}

# Migrate Offer V1
Write-Host "`n=== Migrating Offer V1 (First Edition) ===" -ForegroundColor Cyan
if (Test-Path "offer/SwordRefers_Offer_0000") {
    Move-OfferProblem -SourcePath "offer/SwordRefers_Offer_0000" `
                      -TargetBase "code/interview/sword-offer/offer-v1" `
                      -Version "V1" `
                      -Mapping $offerV1Mapping
}

# Migrate Offer V2
Write-Host "`n=== Migrating Offer V2 (Second Edition) ===" -ForegroundColor Cyan
if (Test-Path "offer/SwordRefers_Offer_II_0000") {
    Move-OfferProblem -SourcePath "offer/SwordRefers_Offer_II_0000" `
                      -TargetBase "code/interview/sword-offer/offer-v2" `
                      -Version "V2" `
                      -Mapping $offerV2Mapping
}

if (Test-Path "offer/SwordRefers_Offer_II_0100") {
    Move-OfferProblem -SourcePath "offer/SwordRefers_Offer_II_0100" `
                      -TargetBase "code/interview/sword-offer/offer-v2" `
                      -Version "V2" `
                      -Mapping $offerV2Mapping
}

Write-Host "`n=== Migration Complete ===" -ForegroundColor Green
Write-Host "New structure:" -ForegroundColor Cyan
Write-Host "  code/interview/sword-offer/offer-v1/  - First Edition"
Write-Host "  code/interview/sword-offer/offer-v2/  - Second Edition"
