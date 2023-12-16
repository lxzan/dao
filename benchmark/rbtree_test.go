package benchmark

//func BenchmarkRBTree_Set(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		var tree = rbtree.New[int, string]()
//		for j := 0; j < bench_count; j++ {
//			tree.Set(j, "")
//		}
//	}
//}
//
//func BenchmarkRBTree_Find(b *testing.B) {
//	var tree = rbtree.New[int, string]()
//	for j := 0; j < bench_count; j++ {
//		tree.Set(j, "")
//	}
//
//	b.StartTimer()
//	for i := 0; i < b.N; i++ {
//		for j := 0; j < bench_count; j++ {
//			tree.Get(j)
//		}
//	}
//}
//
//func BenchmarkRBTree_Delete(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		var tree = rbtree.New[int, string]()
//		for j := 0; j < bench_count; j++ {
//			tree.Set(j, "")
//		}
//
//		for j := 0; j < bench_count; j++ {
//			tree.Delete(j)
//		}
//	}
//}
//
//func BenchmarkRBTree_Between(b *testing.B) {
//	var tree = rbtree.New[int, string]()
//	for j := 0; j < bench_count; j++ {
//		tree.Set(j, "")
//	}
//
//	b.StartTimer()
//	for i := 0; i < b.N; i++ {
//		var left = utils.Rand.Intn(bench_count)
//		var right = left + 10
//		var qb = rbtree.QueryBuilder[int]{
//			LeftFilter:  func(d int) bool { return d >= left },
//			RightFilter: func(d int) bool { return d < right },
//			Limit:       10,
//			Order:       rbtree.DESC,
//		}
//		for j := 0; j < bench_count; j++ {
//			tree.Query(&qb)
//		}
//	}
//	b.StopTimer()
//}
//
//func BenchmarkRBTree_GreaterEqual(b *testing.B) {
//	var tree = rbtree.New[int, string]()
//	for j := 0; j < bench_count; j++ {
//		tree.Set(j, "")
//	}
//
//	b.StartTimer()
//	for i := 0; i < b.N; i++ {
//		var k = utils.Rand.Intn(bench_count)
//		var qb = rbtree.QueryBuilder[int]{
//			LeftFilter: func(d int) bool { return d >= k },
//			Limit:      10,
//			Order:      rbtree.ASC,
//		}
//		for j := 0; j < bench_count; j++ {
//			tree.Query(&qb)
//		}
//	}
//	b.StopTimer()
//}
//
//func BenchmarkRBTree_LessEqual(b *testing.B) {
//	var tree = rbtree.New[int, string]()
//	for j := 0; j < bench_count; j++ {
//		tree.Set(j, "")
//	}
//
//	b.StartTimer()
//	for i := 0; i < b.N; i++ {
//		var k = utils.Rand.Intn(bench_count)
//		var qb = rbtree.QueryBuilder[int]{
//			RightFilter: func(d int) bool { return d <= k },
//			Limit:       10,
//			Order:       rbtree.DESC,
//		}
//		for j := 0; j < bench_count; j++ {
//			tree.Query(&qb)
//		}
//	}
//	b.StopTimer()
//}
//
//func BenchmarkRBTree_GetMinKey(b *testing.B) {
//	var tree = rbtree.New[int, string]()
//	for j := 0; j < bench_count; j++ {
//		tree.Set(j, "")
//	}
//
//	b.StartTimer()
//	for i := 0; i < b.N; i++ {
//		var key = utils.Rand.Intn(bench_count)
//		for j := 0; j < bench_count; j++ {
//			tree.GetMinKey(func(k int) bool {
//				return k >= key
//			})
//		}
//	}
//	b.StopTimer()
//}
//
//func BenchmarkRBTree_GetMaxKey(b *testing.B) {
//	var tree = rbtree.New[int, string]()
//	for j := 0; j < bench_count; j++ {
//		tree.Set(j, "")
//	}
//
//	b.StartTimer()
//	for i := 0; i < b.N; i++ {
//		var key = utils.Rand.Intn(bench_count)
//		for j := 0; j < bench_count; j++ {
//			tree.GetMaxKey(func(k int) bool {
//				return k <= key
//			})
//		}
//	}
//	b.StopTimer()
//}
