function BubbleSort(A)
{
    let n = A.length
    let swapped = true
    while(swapped)
    {
        swapped = false
        for (let i = 0; i < n - 1; i++)
        {
            if (A[i] > A[i+1])
            {
                let temp = A[i]
                A[i] = A[i+1]
                A[i+1] = temp
                swapped = true
            }
        }
    }

    return A
}
