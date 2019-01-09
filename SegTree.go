type SegTree struct {
	st []int
	lazy []int
	has []bool
	size int
	el_neutro int
}

func f(a, b int) int {
	return a+b
}

func NewSegTree(n int) *SegTree {
	t := new(SegTree)
	t.st = make([]int, 4*n)
	t.lazy = make([]int, 4*n)
	t.has = make([]bool, 4*n)
	t.size = n
	return t
}

func (t *SegTree) propagate(sti, stl, str int) {
    if t.has[sti] {
        t.st[sti] = t.lazy[sti]*(str-stl+1)
        if stl != str {
            t.lazy[sti*2 + 1] = t.lazy[sti]
            t.lazy[sti*2 + 2] = t.lazy[sti]

            t.has[sti*2 + 1] = true
            t.has[sti*2 + 2] = true
        }
        t.has[sti] = false
    }
}

func (t *SegTree) query(sti, stl, str, l, r int) int {
    t.propagate(sti, stl, str)

    if str < l || r < stl {
        return t.el_neutro
    }

    if stl >= l && str <= r {
        return t.st[sti]
    }

    mid := (str+stl)/2

    return f(t.query(sti*2+1,stl,mid,l,r),t.query(sti*2+2,mid+1,str,l,r))
}

func (t *SegTree) UpdateRange(sti, stl, str, l, r, amm int) {
    t.propagate(sti, stl, str)
    
    if stl >= l && str <= r {
        t.lazy[sti] = amm
        t.has[sti] = true
        t.propagate(sti, stl, str)
        return
    }

    if stl > r || str < l {
        return
    }
    
    mid := (stl + str)/2
    
    t.UpdateRange(sti*2+1,stl,mid,l,r,amm)
    t.UpdateRange(sti*2+2,mid+1,str,l,r,amm)
    
    t.st[sti] = f(t.st[sti*2+1],t.st[sti*2+2])
}

func (t *SegTree) PQuery(l, r int) int {
    return t.query(0,0,t.size-1,l,r)
}

func (t *SegTree) PUpdateRange(l, r, amm int){
    t.UpdateRange(0,0,t.size-1,l,r,amm);
}