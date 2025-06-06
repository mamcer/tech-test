﻿namespace MergeOrdenado.Cli;

public class Program
{
    public static void Main(string[] args)
    {
        var array1 = new int[] { 1, 3, 5, 7, 9 };
        var array2 = new int[] { 2, 3, 6, 8, 10 };

        Console.WriteLine("array 1: " + string.Join(", ", array1));
        Console.WriteLine("array 2: " + string.Join(", ", array2));
        Console.WriteLine($"merge ordenado result: {string.Join(", ", MergeOrdenado(array1, array2))}");
    }

    /// <summary>
    /// Merges two sorted integer arrays into a single sorted array.
    /// </summary>
    public static int[] MergeOrdenado(int[] array1, int[] array2)
    {
        if (array1 == null)
        {
            throw new ArgumentNullException(nameof(array1));
        }

        if (array2 == null)
        {
            throw new ArgumentNullException(nameof(array2));
        }

        int[] result = new int[array1.Length + array2.Length];
        int i = 0, j = 0, k = 0;

        while (i < array1.Length && j < array2.Length)
        {
            if (array1[i] < array2[j])
            {
                result[k++] = array1[i++];
            }
            else
            {
                result[k++] = array2[j++];
            }
        }

        while (i < array1.Length)
        {
            result[k++] = array1[i++];
        }

        while (j < array2.Length)
        {
            result[k++] = array2[j++];
        }

        return result;
    }
}