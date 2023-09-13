local case = testcase.new()

function case:BeforeRun()
    fromAddr = self.blockchain:GetRandomAccountByGroup()
    local result = self.blockchain:Invoke({
        caller = fromAddr,
        contract = "ERC20", -- contract name is the contract file name under directory invoke/contract
        func = "mint",
        args = {10000000000},
    })
end

function case:Run()
    fromAddr = self.blockchain:GetRandomAccountByGroup()
    toAddr = self.blockchain:GetRandomAccount(fromAddr)
    --print("to addr:" .. toAddr)
    value = self.toolkit.RandInt(1, 100)
    local result = self.blockchain:Invoke({
        caller = fromAddr,
        contract = "ERC20",
        func = "transfer",
        args = {"0x" .. toAddr},
    })
    --self.blockchain:Confirm(result)
    return result
end
return case