# commands

$ go run . add "Lunch" 20
# Expense added successfully (ID: 1)

$ go run . add "Dinner" 10
# Expense added successfully (ID: 2)

$ go run . list
# ID  Date       Description  Amount
# 1   2024-08-06  Lunch        $20
# 2   2024-08-06  Dinner       $10

$ go run . summary
# Total expenses: $30

$ go run . delete 2
# Expense deleted successfully

$ go run . summary
# Total expenses: $20

$ go run . summary  8
# Total expenses for August: $20

https://roadmap.sh/projects/expense-tracker

