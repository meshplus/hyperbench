local case = testcase.new()

function sleep(n)
    os.execute("sleep " .. tonumber(n))
end

function case:BeforeRun()
    tokenAddrList = {"14dC79964da2C08b23698B3D3cc7Ca32193d9955"}
    for i, v in ipairs(tokenAddrList) do
        fromAddr = v
        print("total accounts:" .. self.index.Accounts)
        for i = 1, self.index.Accounts do
            toAddr = self.blockchain:GetRandomAccount(fromAddr)
            self.blockchain:Transfer({
                from = fromAddr, 
                to = toAddr,
                amount = 10000,
                extra = "11",
            })
        end
    end
end

function case:Run()
    --print("----start lua vm-----")
    fromAddr = self.blockchain:GetRandomAccountByGroup()
    toAddr = self.blockchain:GetRandomAccount(fromAddr)
    --print("from addr:" .. fromAddr .. "to addr:" .. toAddr)
    value = self.toolkit.RandInt(1, 100)
    local ret = self.blockchain:Transfer({
        from = fromAddr, --"14dC79964da2C08b23698B3D3cc7Ca32193d9955",
        to = toAddr, --"90f79bf6eb2c4f870365e785982e1f101e93b906",
        amount = value * 0.01,
        extra = "11",
    })
    --print(ret)
    -- while(true)
    -- do
    --     sleep(10)
    --     confirmResult = self.blockchain:Confirm(ret)
    --     if confirmResult.Status == "confirm" then
    --         break
    --     end
    -- end
    return ret
end

return case
