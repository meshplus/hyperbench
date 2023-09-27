local case = testcase.new()

function case:BeforeRun()
    -- -- transfer token
    -- local tokenAddrList = {
    --     "f39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
    --     "70997970C51812dc3A010C7d01b50e0d17dc79C8",
    --     "3C44CdDdB6a900fa2b585dd299e03d12FA4293BC",
    --     "90F79bf6EB2c4f870365E785982E1f101E93b906", 
    --     "15d34AAf54267DB7D7c367839AAf71A00a2C6A65",
    --     "9965507D1a55bcC2695C58ba16FB37d819B0A4dc",
    --     "976EA74026E726554dB657fA54763abd0C3a0aa9",
    --     "14dC79964da2C08b23698B3D3cc7Ca32193d9955",
    --     "23618e81E3f5cdF7f54C3d65f7FBc0aBf5B21E8f",
    --     "a0Ee7A142d267C1f36714E4a8F75612F20a79720",
    --     "Bcd4042DE499D14e55001CcbB24a551F3b954096",
    --     "71bE63f3384f5fb98995898A86B02Fb2426c5788",
    --     "FABB0ac9d68B0B445fB7357272Ff202C5651694a",
    --     "1CBd3b2770909D4e10f157cABC84C7264073C9Ec",
    --     "dF3e18d64BC6A983f673Ab319CCaE4f1a57C7097",
    --     "cd3B766CCDd6AE721141F452C550Ca635964ce71",
    --     "2546BcD3c84621e976D8185a91A922aE77ECEc30",
    --     "bDA5747bFD65F08deb54cb465eB87D40e51B197E",
    --     "dD2FD4581271e230360230F9337D5c0430Bf44C0",
    --     "8626f6940E2eb28930eFb4CeF49B2d1F2C9C1199",
    -- }
    -- local engineNum = self.index.Engine
    -- local tokenLen = #tokenAddrList
    -- if engineNum > tokenLen then
    --     print("please set engine.cap num: " .. engineNum .. "smaller than token list length:" .. tokenLen .. " when call before run")
    --     return
    -- end
    -- local accountNum = self.index.Accounts
    -- local index = self.index.VM
    -- --print("accounts num:" .. self.index.Accounts)
    -- local from = tokenAddrList[index + 1]
    -- local result
    -- for i=13579,accountNum do
    --     local toAddr = self.blockchain:GetAccount(i-1)
    --     if toAddr ~= from then
    --         result = self.blockchain:Transfer({
    --             from = from, 
    --             to = toAddr,
    --             amount = 1000,
    --             extra = "11",
    --         })
    --     end
    -- end

    -- -- wait token confirm
    -- self.blockchain:Confirm(result)

    -- mint erc20
    -- for i=1,13579 do
    --     local index = i % self.index.Engine
    --     if index == self.index.VM then
    --         local fromAddr = self.blockchain:GetAccount(i-1)
    --         result = self.blockchain:Invoke({
    --             caller = fromAddr,
    --             contract = "ERC20", -- contract name is the contract file name under directory invoke/contract
    --             func = "mint",
    --             args = {100000000},
    --         })
    --     end
    -- end

    -- wait token confirm
    --self.blockchain:Confirm(result)

    -- set contract address
    --self.blockchain:SetContext('{"contract_name": "ERC20", "contract_addr": "0xE6E340D132b5f46d1e472DebcD681B2aBc16e57E"}')
end

function case:Run()
    -- invoke erc20 contract
    local fromAddr = self.blockchain:GetRandomAccountByGroup()
    local toAddr = self.blockchain:GetRandomAccount(fromAddr)
    --print("to addr:" .. toAddr)
    local random = self.toolkit.RandInt(0, 2)
    local value = self.toolkit.RandInt(1, 100)
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