package scripts

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// TestCmd 测试
var TestCmd = &cobra.Command{
	Use:   "test",
	Short: "测试",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		err := Test(ctx)
		if err != nil {
			fmt.Println(err)
		}
	},
}

// init 初始化
func init() {
	rootCmd.AddCommand(TestCmd)
	TestCmd.Flags().StringVar(&startDate, "start_date", "", "开始日期")
}

// Test 测试
func Test(ctx context.Context) error {
	if startDate == "" {
		return fmt.Errorf("start_date不能为空")
	}
	t, err := time.ParseInLocation("2006-01-02", startDate, time.Local)
	if err != nil {
		return err
	}
	fmt.Printf("start_date:%s\n", t.Format(time.DateTime))

	return nil
}
