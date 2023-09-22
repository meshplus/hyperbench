local case = testcase.new()

function case:Run()
    -- invoke erc20 contract
    local fromAddr = self.blockchain:GetRandomAccountByGroup()
    local toAddr = self.blockchain:GetRandomAccount(fromAddr)
    --print("to addr:" .. toAddr)
    local random = self.toolkit.RandInt(0, 2)
    local value = 2
    local result
    if random == 0 then
        result = self.blockchain:Invoke({
            caller = fromAddr,
            contract = "ERC20",
            func = "transfer",
            args = {toAddr, value},
        })
    else
        result = self.blockchain:Invoke({
            caller = fromAddr,
            contract = "ERC20",
            func = "approve",
            args = {toAddr, value},
        })

        recvAddr = self.blockchain:GetRandomAccount(fromAddr)
        result = self.blockchain:Invoke({
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