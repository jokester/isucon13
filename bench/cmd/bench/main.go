package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/isucon/isucon13/bench/internal/config"
	"github.com/isucon/isucon13/bench/isupipe"
	"github.com/isucon/isucon13/bench/scenario"
)

func init() {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatalln(err)
	}
	time.Local = loc

	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {
	ctx := context.Background()

	if config.AdvertiseCost < 1 || 10 < config.AdvertiseCost {
		log.Fatalln("-c(ベンチマーク走行中の広告費用) は1~10の中で設定してください")
	}

	client, err := isupipe.NewClient()
	if err != nil {
		log.Fatalln(err)
	}

	if err := scenario.Pretest(ctx, client); err != nil {
		log.Fatalf("Pretest: ベンチマーカの初期テストに失敗しました: %s", err.Error())
	}

	// // 各シーズンシナリオが無事達成された際には、一度cancel()を実行してこれ以上シーズンシナリオが進行しないようにするため、
	// // シーズンシナリオごとにctxを切る
	// // 元のcontextは1分間でDone()を詰めるため、各シーズンにもそれが反映される
	// season1Ctx, season1Cancel := context.WithCancel(benchCtx)
	// defer season1Cancel()

	// log.Println("Season1シナリオ走行開始~既存配信者に対する大量のライブコメント/リアクション~")
	// log.Println("Season1シナリオの達成条件: 200,000の利益を獲得すること")
	// if err := benchmarker.season1(season1Ctx, isupipe.DefaultClientBaseURL); err != nil {
	// 	// 単なるエラーではなく、season1を達成できずに終わっただけなので、スコアを表示する
	// 	log.Println("Season1シナリオの達成条件を満たせませんでした")
	// 	// printBenchmarkResult()
	// 	os.Exit(0)
	// }
	// log.Println("Season1シナリオの達成条件: 200,000の利益を獲得を満たしました")
	// season1Cancel()

	// log.Println("Season2シナリオ走行開始~新人配信者の大量予約+新人配信に対する大量のライブコメント/リアクション~")
	// season2Ctx, season2Cancel := context.WithCancel(benchCtx)
	// defer season2Cancel()
	// if err := benchmarker.season2(season2Ctx, isupipe.DefaultClientBaseURL); err != nil {
	// 	log.Println("Season2シナリオの達成条件を満たせませんでした")
	// 	printBenchmarkResult()
	// 	os.Exit(0)
	// }
	// log.Println("Season2シナリオの達成条件: すべての予約リクエストが成功し、合計400,000の利益を獲得を満たしました")
	// season2Cancel()

	// printBenchmarkResult()
}

// func printBenchmarkResult() {
// 	criticalErrors, ok := bencherror.GetFinalErrorMessages()[bencherror.BenchmarkCriticalError.ErrorCode()]
// 	if ok && len(criticalErrors) == 0 {
// 		for i, c := range criticalErrors {
// 			log.Printf("critical-error[%d]: %s\n", i, c)
// 		}

// 		log.Println("final score ==> 0 (denied)")
// 		os.Exit(0)
// 	}

// 	for key, messages := range bencherror.GetFinalErrorMessages() {
// 		if key == bencherror.BenchmarkCriticalError.ErrorCode() {
// 			continue
// 		}

// 		for i, message := range messages {
// 			log.Printf("%s[%d]: %s\n", key, i, message)
// 		}
// 	}

// 	finalScore := benchscore.GetFinalScore()
// 	finalProfit := benchscore.GetFinalProfit()
// 	finalPenalty := benchscore.GetFinalPenalty()

// 	if finalScore+finalProfit < finalPenalty {
// 		log.Println("final score ==> 0")
// 	} else {
// 		log.Printf("final score ==> %d\n", finalScore+finalProfit-finalPenalty)
// 	}
// }
