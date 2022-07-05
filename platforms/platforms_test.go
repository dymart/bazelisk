package platforms

import "testing"

func TestDarwinFallback(t *testing.T) {
	type args struct {
		machineName string
		version     string
	}
	tests := []struct {
		name                 string
		args                 args
		wantAlterMachineName string
	}{
		{
			name: "before 4.1.0, amd64 do not fallback",
			args: args{
				machineName: "amd64",
				version:     "4.0.1",
			},
			wantAlterMachineName: "amd64",
		},
		{
			name: "since 4.1.0, amd64 do not fallback either",
			args: args{
				machineName: "amd64",
				version:     "4.1.0",
			},
			wantAlterMachineName: "amd64",
		},
		{
			name: "before 4.1.0, arm64 not supported, fallback to amd64 on arm64",
			args: args{
				machineName: "arm64",
				version:     "4.0.1",
			},
			wantAlterMachineName: "amd64",
		},
		{
			name: "since 4.1.0, arm64 supported, do not fallback",
			args: args{
				machineName: "arm64",
				version:     "4.1.0",
			},
			wantAlterMachineName: "arm64",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAlterMachineName := DarwinFallback(tt.args.machineName, tt.args.version)
			if gotAlterMachineName != tt.wantAlterMachineName {
				t.Errorf("DarwinFallback() gotAlterMachineName = %v, want %v", gotAlterMachineName, tt.wantAlterMachineName)
			}
		})
	}
}
