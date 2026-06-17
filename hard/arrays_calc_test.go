package hard

import "testing"

func TestCalcMaxMin(t *testing.T) {
    nums := []int{3, 1, 4, 2}
    max, min := CalcMaxMin(nums)
    if max != 4 || min != 1 {
        t.Fatalf("expected max=4,min=1 got max=%d,min=%d", max, min)
    }
}

func TestCalcRemoveDuplicates(t *testing.T) {
    nums := []int{1, 2, 1, 3, 2}
    unique := CalcRemoveDuplicates(nums)
    m := map[int]bool{}
    for _, v := range unique {
        m[v] = true
    }
    for _, v := range []int{1, 2, 3} {
        if !m[v] {
            t.Fatalf("expected %d in unique result", v)
        }
    }
}
