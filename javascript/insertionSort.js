function InsertionSort(A)
{
    let i = 1
    while(i < A.length)
    {
        let j = i
        while(j > 0 && A[j-1] > A[j])
        {
            let temp = A[j]
            A[j] = A[j-1]
            A[j-1] = temp
            j -= 1
        }
        i += 1
    }

    return A
}
