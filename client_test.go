package powerjob_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/goexl/powerjob"
)

func TestJob(t *testing.T) {
	client := powerjob.New().Worker("https://job.sichuancredit.cn", "powerjob-worker-samples").Password("powerjob123").Build().Build()
	fmt.Println(client.Job().Name("测试").Do(context.Background()))
}
