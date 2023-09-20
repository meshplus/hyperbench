local case = testcase.new()

function case:BeforeRun()
    fromAddr = self.blockchain:GetRandomAccountByGroup()
    self.amountToken = 1000000000000000000000
    amount = self.amountToken
    --print(type(self))
    --print(type(self.blockchain))
    --print(type(self.blockchain.GetContractAddrByName))
    local routerAddr = self.blockchain:GetContractAddrByName("UniswapV2Router02")

    local tokens = {"Token1", "Token2", "Token3"}
    local multipliers = {1, 10, 100}

    for i, token in ipairs(tokens) do
        local mintAmount = amount * multipliers[i]

        local result = self.blockchain:Invoke({
            caller = fromAddr,
            contract = token,
            func = "mint",
            args = {fromAddr, mintAmount},
        })

        local approveRes = self.blockchain:Invoke({
            caller = fromAddr,
            contract = token,
            func = "approve",
            args = {routerAddr, mintAmount}
        })
    end

    deadline = os.time() + 1000
    local routerAddr = self.blockchain:GetContractAddrByName("UniswapV2Router02")
    local token1 = self.blockchain:GetContractAddrByName("Token1")
    local token2 = self.blockchain:GetContractAddrByName("Token2")
    local token3 = self.blockchain:GetContractAddrByName("Token3")

    local token1token2Liquid = self.blockchain:Invoke({
        caller = fromAddr,
        contract = "UniswapV2Router02",
        func = "addLiquidity",
        args ={
            token1,
            token2,
            amount * 0.5,
            amount * 0.5 * 10,
            amount * 0.4,
            amount * 0.4 * 10,
            fromAddr,
            deadline,
        },
    })

    local token2token3Liquid = self.blockchain:Invoke({
        caller = fromAddr,
        contract = "UniswapV2Router02",
        func = "addLiquidity",
        args ={
            token2,
            token3,
            amount * 0.5 * 10,
            amount * 0.5 * 100,
            amount * 0.4 * 10,
            amount * 0.4 * 100,
            fromAddr,
            deadline,
        },
    })

    local token1token3Liquid = self.blockchain:Invoke({
        caller = fromAddr,
        contract = "UniswapV2Router02",
        func = "addLiquidity",
        args ={
            token1,
            token3,
            amount * 0.5,
            amount * 0.5 * 100,
            amount * 0.4,
            amount * 0.4 * 100,
            fromAddr,
            deadline,
        },
    })

    return {type = "default addLiquidity", result = token1token3Liquid}
end

function case:Run()
    -- 获取随机地址
    addr = self.blockchain:GetRandomAccountByGroup()
    amount = self.amountToken

    -- 获取一个随机数
    random = self.toolkit.RandInt(0, 1)

    if random == 0 then
        -- 随机选择token对进行添加流动性
        pair_choice = self.toolkit.RandInt(1, 3)
        local tokenA, tokenB
        if pair_choice == 1 then
            tokenA = "Token1"
            tokenB = "Token2"
        elseif pair_choice == 2 then
            tokenA = "Token2"
            tokenB = "Token3"
        else
            tokenA = "Token1"
            tokenB = "Token3"
        end

        -- 授权 TokenA
        local approveTokenA = self.blockchain:Invoke({
            caller = addr,
            contract = tokenA,
            func = "approve",
            args = {self.blockchain:GetContractAddrByName("UniswapV2Router02"), amount * 1}
        })

        -- 授权 TokenB
        local approveTokenB = self.blockchain:Invoke({
            caller = addr,
            contract = tokenB,
            func = "approve",
            args = {self.blockchain:GetContractAddrByName("UniswapV2Router02"), amount * 1}
        })

        -- 添加流动性
        deadline = os.time() + 1000
        local addLiquidity = self.blockchain:Invoke({
            caller = addr,
            contract = "UniswapV2Router02",
            func = "addLiquidity",
            args ={
                self.blockchain:GetContractAddrByName(tokenA),
                self.blockchain:GetContractAddrByName(tokenB),
                amount * 1,
                amount * 1,
                amount * 0.01,
                amount * 0.01,
                addr,
                deadline,
            },
        })
        --print(addLiquidity)
        return {type = "addLiquidity", result = addLiquidity}

    else
        -- mint
        token_choice = self.toolkit.RandInt(1, 3)
        if token_choice == 1 then
            token_name = "Token1"
        elseif token_choice == 2 then
            token_name = "Token2"
        else
            token_name = "Token3"
        end

        local mintResult = self.blockchain:Invoke({
            caller = addr,
            contract = token_name,
            func = "mint",
            args = {addr, amount},
        })

        -- swap
        local swapTokens = {"Token1", "Token2", "Token3"}
        table.remove(swapTokens, token_choice)  -- 移除已经mint的token

        local inputToken = self.blockchain:GetContractAddrByName(swapTokens[1])
        local outputToken = self.blockchain:GetContractAddrByName(swapTokens[2])

        local swapResult = self.blockchain:Invoke({
            caller = addr,
            contract = "UniswapV2Router02",
            func = "swapExactTokensForTokens",
            args = {
                amount * 0.05,  -- 交换的数量
                0,  -- 最小输出数量
                {inputToken, outputToken},  -- 交换路径
                addr,  -- 接收地址
                deadline,
            },
        })
        --print(swapResult)
        return {type = "swapResult", result = swapResult}

    end
end
return case
