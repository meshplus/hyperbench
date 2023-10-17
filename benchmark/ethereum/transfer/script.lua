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
                amount = 1000,
                extra = "11",
            })
        end
    end

    -- wait token confirm
    self.blockchain:Confirm(result)
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
    return ret
end

return case
