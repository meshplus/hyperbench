local case = testcase.new()

function case:Run()
    local ret = case.blockchain:Transfer({
        from = "0x01eeC173917c429901b41b98Ac3DD300e060e698",
        to = "a6df6489927a9d0172185efe68de1f9aace82639",
        amount = 100,
        extra = "",
    })
    return ret
end

return case
