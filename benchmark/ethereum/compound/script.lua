local case = testcase.new()


-- Aries
function case:BeforeRun()
   -- set contract address
   self.blockchain:SetContext('{"contract_name": "UNI", "contract_addr": "0x9dfDf098804aEAD808aA5896EAcd0475CCE18409"}')
   self.blockchain:SetContext('{"contract_name": "USDC", "contract_addr": "0x9cB42e0D49fffA7381a0509b0c8909D125c09912"}')
   self.blockchain:SetContext('{"contract_name": "CUNI", "contract_addr": "0x8E5ea1F851AA25f16C77161ba682480c959dC76D"}')
   self.blockchain:SetContext('{"contract_name": "CUSDC", "contract_addr": "0x8f9fe2A462e1a8024541Ae5829821Fd72798cB96"}')
   self.blockchain:SetContext('{"contract_name": "Comptroller", "contract_addr": "0xC7B977C41aE7B25997263b9DF21C17694E6D8429"}')

   
   local fromAddr = self.blockchain:GetRandomAccountByGroup()
   print("from addr:" .. fromAddr)
   local uniMint=100000000000000000000
   local usdcMint=100000000000000000000
   local cUNIAddr = self.blockchain:GetContractAddrByName("CUNI")
   local cUSDCAddr = self.blockchain:GetContractAddrByName("CUSDC")

   self.blockchain:Invoke({
    caller = fromAddr,
    contract = "UNI",
    func = "mint",
    args = {fromAddr, uniMint},
   })

   self.blockchain:Invoke({
      caller = fromAddr,
      contract = "UNI",
      func = "approve",
      args = {cUNIAddr, uniMint},
   })

   self.blockchain:Invoke({
      caller = fromAddr,
      contract = "USDC",
      func = "mint",
      args = {fromAddr, usdcMint},
   })

   self.blockchain:Invoke({
      caller = fromAddr,
      contract = "USDC",
      func = "approve",
      args = {cUSDCAddr, uniMint},
   })
   
   self.blockchain:Invoke({
    caller = fromAddr,
    contract = "Comptroller",
    func = "enterOneMarkets",
    args = {cUSDCAddr},
   })
   self.blockchain:Invoke({
      caller = fromAddr,
      contract = "Comptroller",
      func = "enterOneMarkets",
      args = {cUNIAddr},
   })

   self.blockchain:Invoke({
      caller = fromAddr,
      contract = "CUNI",
      func = "mint",
      args = {uniMint},
   })

   self.blockchain:Invoke({
      caller = fromAddr,
      contract = "CUSDC",
      func = "mint",
      args = {usdcMint},
   })


end


function case:Run()
   local fromAddr = self.blockchain:GetRandomAccountByGroup()
   print("from addr:" .. fromAddr)
   local cUNIAddr = self.blockchain:GetContractAddrByName("CUNI")
   local cUSDCAddr = self.blockchain:GetContractAddrByName("CUSDC")
   local ctokenArray = {UNIAddr,cUSDCAddr}
   local fromAddr = self.blockchain:GetRandomAccountByGroup()
   local uniMint=10000000000000000000

   self.blockchain:Invoke({
      caller = fromAddr,
      contract = "UNI",
      func = "mint",
      args = {fromAddr, uniMint},
   })


   self.blockchain:Invoke({
      caller = fromAddr,
      contract = "UNI",
      func = "approve",
      args = {cUNIAddr, uniMint},
   })

   self.blockchain:Invoke({
      caller = fromAddr,
      contract = "CUNI",
      func = "mint",
      args = {uniMint},
   })

   borrowNum=self.toolkit.RandInt(100000, 1000000)
   self.blockchain:Invoke({
      caller = fromAddr,
      contract = "CUSDC",
      func = "borrow",
      args = {borrowNum},
   })


   self.blockchain:Invoke({
      caller = fromAddr,
      contract = "USDC",
      func = "approve",
      args = {cUSDCAddr,borrowNum},
   })

   self.blockchain:Invoke({
      caller = fromAddr,
      contract = "CUSDC",
      func = "repayBorrow",
      args = {borrowNum},
   })

   

   







end
return case
