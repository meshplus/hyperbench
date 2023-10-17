local case = testcase.new()


-- Aries
function case:BeforeRun()
   -- set contract address
   self.blockchain:SetContext('{"contract_name": "UNI", "contract_addr": "0x58c88Ae044A5471CC90472bCe34b67a7432Df716"}')
   self.blockchain:SetContext('{"contract_name": "USDC", "contract_addr": "0x6d1448CeC252968f2E2526f144C7eedf1cb141e5"}')
   self.blockchain:SetContext('{"contract_name": "CUNI", "contract_addr": "0x12b7358f0B1e2874C6114ecFa3Dc73a3a731272F"}')
   self.blockchain:SetContext('{"contract_name": "CUSDC", "contract_addr": "0xb0aCfACcE6946eEfE99A774bfb14d44bd276dAca"}')
   self.blockchain:SetContext('{"contract_name": "Comptroller", "contract_addr": "0xA853A791361D00b029baf1efB856578D1C681813"}')

   
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
   local cUNIAddr = self.blockchain:GetContractAddrByName("CUNI")
   local cUSDCAddr = self.blockchain:GetContractAddrByName("CUSDC")
   local ctokenArray = {UNIAddr,cUSDCAddr}
   local fromAddr = self.blockchain:GetRandomAccountByGroup()
   local uniMint=10000000000000000000
   local mintNum=10000000000000000

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
      contract = "UNI",
      func = "mint",
      args = {fromAddr, uniMint},
   })


   self.blockchain:Invoke({
      caller = fromAddr,
      contract = "UNI",
      func = "approve",
      args = {cUNIAddr, mintNum},
   })

   self.blockchain:Invoke({
      caller = fromAddr,
      contract = "CUNI",
      func = "mint",
      args = {mintNum},
   })

   borrowNum=self.toolkit.RandInt(100000, 1000000)
   self.blockchain:Invoke({
      caller = fromAddr,
      contract = "CUNI",
      func = "borrow",
      args = {borrowNum},
   })


   self.blockchain:Invoke({
      caller = fromAddr,
      contract = "UNI",
      func = "approve",
      args = {cUNIAddr,borrowNum},
   })

   self.blockchain:Invoke({
      caller = fromAddr,
      contract = "CUNI",
      func = "repayBorrow",
      args = {borrowNum},
   })

   

   







end
return case
