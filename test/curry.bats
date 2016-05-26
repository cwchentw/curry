#!/usr/bin/env bats

@test "update" {
    run curry update
    [ "$status" -eq 0 ]
}

@test "list" {
    run curry list
    [ "$status" -eq 0 ]
    [ "$output" != "" ]
}

@test "favor" {
    run curry favor TWD
    [ "$status" -eq 0 ]
}

@test "from to same currency" {
    run curry from TWD to TWD 1
    [ "$status" -eq 0 ]
    [ "$output" = "1.00" ]
}

@test "from to different currencies" {
    run curry from EUR to USD 1
    [ "$status" -eq 0 ]
    [ "$output" != "1.00" ]
}