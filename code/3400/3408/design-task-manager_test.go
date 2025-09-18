package _3408

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	type args struct {
		tasks [][]int
	}
	tests := []struct {
		name string
		args args
		want TaskManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.tasks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskManager_Add(t *testing.T) {
	type args struct {
		userId   int
		taskId   int
		priority int
	}
	tests := []struct {
		name string
		this *TaskManager
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.Add(tt.args.userId, tt.args.taskId, tt.args.priority)
		})
	}
}

func TestTaskManager_Edit(t *testing.T) {
	type args struct {
		taskId      int
		newPriority int
	}
	tests := []struct {
		name string
		this *TaskManager
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.Edit(tt.args.taskId, tt.args.newPriority)
		})
	}
}

func TestTaskManager_Rmv(t *testing.T) {
	type args struct {
		taskId int
	}
	tests := []struct {
		name string
		this *TaskManager
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.Rmv(tt.args.taskId)
		})
	}
}

func TestTaskManager_ExecTop(t *testing.T) {
	tests := []struct {
		name string
		this *TaskManager
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.ExecTop(); got != tt.want {
				t.Errorf("TaskManager.ExecTop() = %v, want %v", got, tt.want)
			}
		})
	}
}
