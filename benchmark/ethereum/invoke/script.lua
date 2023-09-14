local case = testcase.new()

function case:BeforeRun()
    for i = 0, 100000 do
        fromAddr = self.blockchain:GetRandomAccountByGroup()
        local result = self.blockchain:Invoke({
            caller = fromAddr,
            contract = "ERC20", -- contract name is the contract file name under directory invoke/contract
            func = "mint",
            args = {10000000000},
        })
        --print("before call result: " .. result.UID)
    end
end

function case:Run()
    fromAddr = self.blockchain:GetRandomAccountByGroup()
    toAddr = self.blockchain:GetRandomAccount(fromAddr)
    --print("to addr:" .. toAddr)
    random = self.toolkit.RandInt(0, 2)
    value = self.toolkit.RandInt(1, 100)
    if random == 0 then
        local result = self.blockchain:Invoke({
            caller = fromAddr,
            contract = "ERC20",
            func = "transfer",
            args = {toAddr, value},
        })
    else
        local result = self.blockchain:Invoke({
            caller = fromAddr,
            contract = "ERC20",
            func = "approve",
            args = {toAddr, value},
        })

        recvAddr = self.blockchain:GetRandomAccount(fromAddr)
        local result = self.blockchain:Invoke({
            caller = toAddr,
            contract = "ERC20",
            func = "transferFrom",
            args = {fromAddr, recvAddr, value},
        })
    end

    --print("call result:" .. result.UID)
    --self.blockchain:Confirm(result)
    return result
end
return case