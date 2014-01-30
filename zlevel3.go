package blas

func Zgemm(tA, tB Transpose, alpha complex128, A, B GeneralCmplx, beta complex128, C GeneralCmplx) {
	var m, n, k int
	if tA == NoTrans {
		m, k = A.Rows, A.Cols
	} else {
		m, k = A.Cols, A.Rows
	}
	if tB == NoTrans {
		n = B.Cols
		if k != B.Rows {
			panic("blas: dimension mismatch")
		}
	} else {
		n = B.Rows
		if k != B.Cols {
			panic("blas: dimension mismatch")
		}
	}
	if m != C.Rows {
		panic("blas: dimension mismatch")
	}
	if n != C.Cols {
		panic("blas: dimension mismatch")
	}
	implCmplx.Zgemm(A.Order, tA, tB, m, n, k, alpha, A.Data, A.Stride,
		B.Data, B.Stride, beta, C.Data, C.Stride)
}

func Zsymm(s Side, alpha complex128, A SymmetricCmplx, B GeneralCmplx, beta complex128, C GeneralCmplx) {
	var m, n int
	if s == Left {
		m = A.N
		n = B.Cols
		if m != B.Rows {
			panic("blas: dimension mismatch")
		}
	} else {
		m = B.Rows
		n = A.N
		if n != B.Cols {
			panic("blas: dimension mismatch")
		}
	}
	if m != C.Rows {
		panic("blas: dimension mismatch")
	}
	if n != C.Cols {
		panic("blas: dimension mismatch")
	}
	implCmplx.Zsymm(A.Order, s, A.Uplo, m, n, alpha, A.Data, A.Stride,
		B.Data, B.Stride, beta, C.Data, C.Stride)
}

func Zsyrk(t Transpose, alpha complex128, A GeneralCmplx, beta complex128, C SymmetricCmplx) {
	var n, k int
	if t == NoTrans {
		n, k = A.Rows, A.Cols
	} else {
		n, k = A.Cols, A.Rows
	}
	if n != C.N {
		panic("blas: dimension mismatch")
	}
	implCmplx.Zsyrk(A.Order, C.Uplo, t, n, k, alpha, A.Data, A.Stride, beta,
		C.Data, C.Stride)
}

func Zsyr2k(t Transpose, alpha complex128, A, B GeneralCmplx, beta complex128, C SymmetricCmplx) {
	var n, k int
	if t == NoTrans {
		n, k = A.Rows, A.Cols
		if n != B.Rows || k != B.Cols {
			panic("blas: dimension mismatch")
		}
	} else {
		n, k = A.Cols, A.Rows
		if k != B.Rows || n != B.Cols {
			panic("blas: dimension mismatch")
		}
	}
	if n != C.N {
		panic("blas: dimension mismatch")
	}
	implCmplx.Zsyr2k(A.Order, C.Uplo, t, n, k, alpha, A.Data, A.Stride,
		B.Data, B.Stride, beta, C.Data, C.Stride)
}

func Ztrmm(s Side, tA Transpose, alpha complex128, A TriangularCmplx, B GeneralCmplx) {
	if s == Left {
		if A.N != B.Rows {
			panic("blas: dimension mismatch")
		}
	} else {
		if A.N != B.Cols {
			panic("blas: dimension mismatch")
		}
	}
	implCmplx.Ztrmm(A.Order, s, A.Uplo, tA, A.Diag, B.Rows, B.Cols, alpha, A.Data, A.Stride,
		B.Data, B.Stride)
}

func Ztrsm(s Side, tA Transpose, alpha complex128, A TriangularCmplx, B GeneralCmplx) {
	if s == Left {
		if A.N != B.Rows {
			panic("blas: dimension mismatch")
		}
	} else {
		if A.N != B.Cols {
			panic("blas: dimension mismatch")
		}
	}
	implCmplx.Ztrsm(A.Order, s, A.Uplo, tA, A.Diag, B.Rows, B.Cols, alpha, A.Data, A.Stride,
		B.Data, B.Stride)
}

func Zhemm(s Side, alpha complex128, A Hermitian, B GeneralCmplx, beta complex128, C GeneralCmplx) {
	var m, n int
	if s == Left {
		m = A.N
		n = B.Cols
		if m != B.Rows {
			panic("blas: dimension mismatch")
		}
	} else {
		m = B.Rows
		n = A.N
		if n != B.Cols {
			panic("blas: dimension mismatch")
		}
	}
	if m != C.Rows {
		panic("blas: dimension mismatch")
	}
	if n != C.Cols {
		panic("blas: dimension mismatch")
	}
	implCmplx.Zhemm(A.Order, s, A.Uplo, m, n, alpha, A.Data, A.Stride,
		B.Data, B.Stride, beta, C.Data, C.Stride)
}

func Zherk(t Transpose, alpha float64, A GeneralCmplx, beta float64, C Hermitian) {
	var n, k int
	if t == NoTrans {
		n, k = A.Rows, A.Cols
	} else {
		n, k = A.Cols, A.Rows
	}
	if n != C.N {
		panic("blas: dimension mismatch")
	}
	implCmplx.Zherk(A.Order, C.Uplo, t, n, k, alpha, A.Data, A.Stride, beta,
		C.Data, C.Stride)
}

func Zher2k(t Transpose, alpha complex128, A, B GeneralCmplx, beta float64, C Hermitian) {
	var n, k int
	if t == NoTrans {
		n, k = A.Rows, A.Cols
		if n != B.Rows || k != B.Cols {
			panic("blas: dimension mismatch")
		}
	} else {
		n, k = A.Cols, A.Rows
		if k != B.Rows || n != B.Cols {
			panic("blas: dimension mismatch")
		}
	}
	if n != C.N {
		panic("blas: dimension mismatch")
	}
	implCmplx.Zher2k(A.Order, C.Uplo, t, n, k, alpha, A.Data, A.Stride,
		B.Data, B.Stride, beta, C.Data, C.Stride)
}
