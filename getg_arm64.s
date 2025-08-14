// +build go1.23,!go1.26

// Copyright (c) 2024 Outrig, Inc.
// Licensed under the Apache License, Version 2.0

#include "textflag.h"

// func getg() *g
TEXT ·getg(SB),NOSPLIT,$0-8
    MOVD g, R0
    MOVD R0, ret+0(FP)
    RET

