local case = testcase.new()

function case:BeforeRun()
    -- set contract addresses
    self.blockchain.SetContext('{"contract_name": "aave", "contract_addr": "0x8811817d4982fC4Ea24ecAA7e7Ec502069b0a353"}')
    self.blockchain.SetContext('{"contract_name": "lendingpool", "contract_addr": "0xC6E282fbeC6aFe4EAC7a6A41369b519E7a9209cd"}')
end

function case:Run()
    -- invoke erc20 contract to approve
    local fromAddr = self.blockchain:GetAccount(0)
    local anotherAddr = self.blockchain:GetAccount(1)
    local lendingPoolAddr = self.blockchain:GetContractAddrByName("lendingpool")
    print("to addr:" .. lendingPoolAddr)
    local value = self.toolkit.RandInt(1, 100)
    local approveRes = self.blockchain:Invoke({
        caller = fromAddr,
        contract = "aave",
        func = "approve",
        args = {lendingPoolAddr, value},
    })
    print("approve aave result:" .. approveRes.UID)

    -- deposit some aave to lendingPool
    local aaveAddr = self.blockchain:GetContractAddrByName("aave")
    local depositRes = self.blockchain.Invoke({
        caller = fromAddr,
        contract = "lendingpool",
        func = "deposit",
        args = {aaveAddr, value, fromAddr, 0},
    })
    print("deposit aave result:" .. depositRes.UID)

    -- withdraw some aave to another account
    local withdrawRes = self.blockchain.Invoke({
        caller = fromAddr,
        contract = "lendingpool",
        func = "withdraw",
        args = {aaveAddr, value, anotherAddr},
    })
    print("withdraw aave result:" .. withdrawRes.UID)
    return withdrawRes
end
return case