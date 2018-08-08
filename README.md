# TaskCouponD
Akhil runs a company. He wants to give salary hike to employees in his company. He writes all
the employee names in a sheet of paper. The order of names on the sheet is fixed (their positions
are fixed), and each of them has a rating score according to his or her performance in the
company. Akhil wants to give at least 1 rupee hike to each employee. If two employees are
positioned next to each other, then the one with the higher rating must get more hike. Akhil also
wants to save money, so he needs to minimize the total number of amount given to the
employees.

Input format: the first line of the input is an integer N, the number of employees in Akhil’s
company. Each of the following N lines contains an integer that indicates the rating of each
employee.

Constraints:

1 <= N <= 10^5

1 <= rating <= 10^5

Output format: Output a single line containing the minimum number of amount Akhil would
spend in total.

Note: when two children have equal rating, they are allowed to have different amount of hike.

Sample input:

3

1

2

2

Sample output:

4

Explanation:

Here 1, 2, 2 is the rating. Note that when two employees have equal rating, they are allowed to
have different amount in hike. Hence optimal distribution will be 1, 2, 1. deals and makes online
shopping more rewarding.

Sample input:
10

2

4

2

6

1

7

8

9

2

1

Sample output:

19

Explanation:

This would be the amount for each employee

1, 2, 1, 2, 1, 2, 3, 4, 2, 1


# Installation 
Golang Installation Link
https://manishpaneri.blogspot.com/2018/03/how-to-install-go-19-on-ubuntu.html

# Set Env Variable
export OUTPUT_PATH=/home/testcaseinput
