 
# This line create a module in go
go mod init firstmodule
# after this command you should see the
# message below and find the file go.mod
# in the same folder
#
# go: creating new go.mod: module firstmodule
# go: to add module requirements and sums:
#        go mod tidy
#

# After you crate a module, you 
# can also run go build to compile
# the code

# it will get all code in its folder
# and generate an exec file with the
# module name
go build

# go install buils the file to go installation folder
# I still have to learn more about it
go install

# the command below import a package into go.mod
# go get 
go get github.com/badoux/checkmail

# Case there are unused packages, you can run
# the below command to remove them of go.mod file 
go mod tidy