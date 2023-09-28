local case = testcase.new()

function sleep(n)
    os.execute("sleep " .. tonumber(n))
end

function case:BeforeRun()
    -- transfer token
    local tokenAddrList = {
        "f39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
        "70997970C51812dc3A010C7d01b50e0d17dc79C8",
        "3C44CdDdB6a900fa2b585dd299e03d12FA4293BC",
        "90F79bf6EB2c4f870365E785982E1f101E93b906",
        "15d34AAf54267DB7D7c367839AAf71A00a2C6A65",
        "9965507D1a55bcC2695C58ba16FB37d819B0A4dc",
        "976EA74026E726554dB657fA54763abd0C3a0aa9",
        "14dC79964da2C08b23698B3D3cc7Ca32193d9955",
        "23618e81E3f5cdF7f54C3d65f7FBc0aBf5B21E8f",
        "a0Ee7A142d267C1f36714E4a8F75612F20a79720",
        "Bcd4042DE499D14e55001CcbB24a551F3b954096",
        "71bE63f3384f5fb98995898A86B02Fb2426c5788",
        "FABB0ac9d68B0B445fB7357272Ff202C5651694a",
        "1CBd3b2770909D4e10f157cABC84C7264073C9Ec",
        "dF3e18d64BC6A983f673Ab319CCaE4f1a57C7097",
        "cd3B766CCDd6AE721141F452C550Ca635964ce71",
        "2546BcD3c84621e976D8185a91A922aE77ECEc30",
        "bDA5747bFD65F08deb54cb465eB87D40e51B197E",
        "dD2FD4581271e230360230F9337D5c0430Bf44C0",
        "8626f6940E2eb28930eFb4CeF49B2d1F2C9C1199",
    }
    local engineNum = self.index.Engine
    local tokenLen = #tokenAddrList
    if engineNum > tokenLen then
        print("please set engine.cap num: " .. engineNum .. "smaller than token list length:" .. tokenLen .. " when call before run")
        return
    end
    local accountNum = self.index.Accounts
    local index = self.index.VM
    --print("accounts num:" .. self.index.Accounts)
    local from = tokenAddrList[index + 1]
    local result
    for i=1,accountNum do
        local toAddr = self.blockchain:GetAccount(i-1)
        if toAddr ~= from then
            result = self.blockchain:Transfer({
                from = from,
                to = toAddr,
                amount = 100000,
                extra = "11",
            })
        end
        sleep(0.1)
    end

    -- wait token confirm
    self.blockchain:Confirm(result)

    self.blockchain.SetContext('{"contract_name": "UniswapV2Router02", "contract_addr": "0xb272962cCFd20F705cB7639dd0b26C6981D5b9bd"}')
    fromAddr = self.blockchain:GetRandomAccountByGroup()
    self.amountToken = 1000000000000000000000
    amount = self.amountToken
    self.routerAddr = self.blockchain:GetContractAddrByName("UniswapV2Router02")
    print(self.routerAddr)

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
            args = {self.routerAddr, mintAmount}
        })
    end

    deadline = os.time() + 10000
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

    -- wait confirm
    self.blockchain:Confirm(token1token2Liquid)

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

    -- wait confirm
    self.blockchain:Confirm(token2token3Liquid)

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

    -- wait confirm
    self.blockchain:Confirm(token1token3Liquid)

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
            args = {self.routerAddr, amount * 1}
        })

        -- wait confirm
        self.blockchain:Confirm(approveTokenA)

        -- 授权 TokenB
        local approveTokenB = self.blockchain:Invoke({
            caller = addr,
            contract = tokenB,
            func = "approve",
            args = {self.routerAddr, amount * 1}
        })
        print("tx of approve: ", approveTokenB.UID)

        -- wait confirm
        self.blockchain:Confirm(approveTokenB)

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
                1000,
                1000,
                addr,
                deadline,
            },
        })
        -- wait confirm
        self.blockchain:Confirm(addLiquidity)

        print("tx of addLiquidity: ", addLiquidity.UID)
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

        -- wait confirm
        self.blockchain:Confirm(mintResult)

        local approveToken = self.blockchain:Invoke({
            caller = addr,
            contract = token_name,
            func = "approve",
            args = {self.routerAddr, amount}
        })

        -- wait confirm
        self.blockchain:Confirm(approveToken)

        -- swap
        local swapTokens = {"Token1", "Token2", "Token3"}
        table.remove(swapTokens, token_choice)  -- 移除已经mint的token

        local inputToken = self.blockchain:GetContractAddrByName(swapTokens[1])
        local outputToken = self.blockchain:GetContractAddrByName(swapTokens[2])

        local swapResult = self.blockchain:Invoke({
            caller = addr,
            contract = "UniswapV2Router02",
            func = "swapExactTokensForTokensSupportingFeeOnTransferTokens",
            args = {
                amount * 0.05,  -- 交换的数量
                0,  -- 最小输出数量
                {inputToken, outputToken},  -- 交换路径
                addr,  -- 接收地址
                deadline,
            },
        })

        -- wait confirm
        self.blockchain:Confirm(swapResult)

        return {type = "swapResult", result = swapResult}

    end
end
return case
