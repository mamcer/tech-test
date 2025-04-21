#include <stdio.h>

int maximumDifference(int a[], int asize)
{
    int md = -1;
    for (int i = 0; i < asize; i++)
    {
        for (int j = i + 1; j < asize; j++)
        {
            if (a[i] < a[j])
            {
                int cd = a[j] - a[i];
                if (cd > md)
                {
                    md = cd;
                }
            }
        }
    }
    return md;
}

int main()
{
    int arr[] = {15, 3, 6, 10};
    int maxDifference = maximumDifference(arr, sizeof(arr)/sizeof(int));
    printf("maximum difference: %d\n", maxDifference);
}