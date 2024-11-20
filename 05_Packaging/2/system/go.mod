module github.com/CrisPassos/goexpert/5_Packaging/2/system

go 1.23.2

// RUN go mod edit -replace github.com/CrisPassos/goexpert/5_Packaging/2/math=../math
replace github.com/CrisPassos/goexpert/5_Packaging/2/math => ../math

require github.com/CrisPassos/goexpert/5_Packaging/2/math v0.0.0-00010101000000-000000000000
