local case = testcase.new()

function case:BeforeRun()
    -- axiom aries
    --self.blockchain.SetContext('{"contract_name": "3pool", "contract_addr": "0x34737E9A2606B34707e4E66a1ccFfb3aDB759f58"}')

    -- 131 server
    self.blockchain.SetContext('{"contract_name": "3pool", "contract_addr": "0x8645A21a317F7ec4354124ae11597Bd2C0C214F0"}')
end

function case:Run()
    local fromAddr = self.blockchain:GetAccount(0)
    -- liquidity for coin1, coin2 and coin3
    local amount1 = self.toolkit.RandInt(1, 100)
    local amount2 = self.toolkit.RandInt(1, 100)
    local amount3 = self.toolkit.RandInt(1, 100)
    -- lp token amount for mint and burn
    local lpTokenAmount = self.toolkit.RandInt(1, 100)
    -- add liquidity for 3 coins
    res = self.blockchain:Invoke({
        caller = fromAddr,
        contract = "3pool",
        func = "add_liquidity",
        args = {
            amount1, amount2, amount3, lpTokenAmount
        },
    })
    --self.blockchain:Confirm(res)
    -- calculate amount of burn lp token before remove liquidity
    tokenAmount = self.blockchain.Invoke({
        caller = fromAddr,
        contract = "3pool",
        func = "calc_token_amount",
        args = {
            amount1, amount2, amount3, false,
        },
    })
    --self.blockchain:Confirm(tokenAmount)
    -- remove liquidity for 3 coins
    res=self.blockchain.Invoke({
        caller = fromAddr,
        contract = "3pool",
        func = "remove_liquidity",
        args = {
            lpTokenAmount, amount1, amount2, amount3
        },
    })
    return res
end
return case