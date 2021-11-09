local case = testcase.new()
function case:Run()
    local result = self.blockchain:Invoke({
        func = "",
        args = {""},
    })
    return result
end
return case