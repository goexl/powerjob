package powerjob_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/goexl/powerjob"
)

func TestJob(t *testing.T) {
	client := powerjob.New().Worker("https://job.sichuancredit.cn", "powerjob-worker-samples").Password("powerjob123").Build().Build()
	var id int64
	fmt.Println(client.Job().Id(&id).Http("https://www.baidu.com").Build().Do(context.Background()))
	fmt.Println(id)
}
