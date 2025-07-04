// +build go1.23,!go1.25

// Copyright (c) 2024 Outrig, Inc.
// Licensed under the Apache License, Version 2.0

#include "textflag.h"

// func getg() *g
TEXT ·getg(SB),NOSPLIT,$0-8
    MOVQ (TLS), R14
    MOVQ R14, ret+0(FP)
    RET
