namespace Maximum.Cli;

using System;

public class Program
{
    public static void Main(string[] args)
    {
        int[] input = { 15, 3, 6, 10 };
        Console.WriteLine("the maximum difference is: " + MaximumDifference(input));
    }

    public static int MaximumDifference(int[] arr)
    {
        int maxDiff = -1;

        for (int i = 0; i < arr.Length; i++)
        {
            for (int j = i + 1; j < arr.Length; j++)
            {
                if (arr[j] > arr[i])
                {
                    if ( (arr[j] - arr[i]) > maxDiff)
                    {
                        maxDiff = arr[j] - arr[i];
                    }
                }
            }
        }

        return maxDiff;
    }
}