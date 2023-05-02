package config
import (
"os"
"testing"
"time"
)

func Test_getEnvDuration(t *testing.T) {
	type args struct{
		key string
		defaultValue string
	}
	tests := []struct {
		name string
		desc string
		args args
		want time.Duration
	}{
		{
			name: "test1",
			desc: "needs to be 60 seconds",
			args: args{
				key: "CACHE_DEFAULT_EXPIRATION",
				defaultValue: "60s",
			},
			want: 60 * time.Second,
			},
		{
			name: "test2",
			desc: "needs to be 0 seconds",
			args: args{
				key: "CACHE_DEFAULT_EXPIRATION",
				defaultValue: "wrong value",
			},
			want: 0 * time.Second,
			},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getEnvDuration(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("getEnvDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getEnvDurationWithEnv(t *testing.T) {
	want := 30 * time.Second 
	os.Setenv("CACHE_DEFAULT_EXPIRATION", "30s")
	
	value := getEnvDuration("CACHE_DEFAULT_EXPIRATION", "60s")
	if value != want{
		t.Errorf("getEnvDuration() = %v, want %v", value, want)
	}
	os.Unsetenv("CACHE_DEFAULT_EXPIRATION")
}
