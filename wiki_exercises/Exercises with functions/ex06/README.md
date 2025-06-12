Write a program that uses the valuePayment function to determine the amount to be paid for an installment of a bill.

The program should ask the user for the installment amount and the number of days late and pass these values ​​to the valuePayment function, which will calculate the amount to be paid and return this value to the program that called it.

The program should then display the amount to be paid on the screen. After execution, the program should ask for another installment amount again and continue in this way until a value equal to zero is entered for the installment.

At this point, the program should be closed, displaying the report for the day, which will contain the number and total amount of installments paid on the day.

The calculation of the amount to be paid is done as follows. For payments without delay, charge the installment amount.

When there is a delay, charge a 3% fine, plus 0.1% interest per day late.