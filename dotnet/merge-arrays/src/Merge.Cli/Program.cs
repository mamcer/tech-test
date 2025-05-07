namespace Merge.Cli;

public class Program
{
    public static void Main(string[] args)
    {
        var array01 = new[] { 1, 3, 5, 7, 9 };
        var array02 = new[] { 2, 4, 6, 8, 10 };
        var mergedArray = MergeArrays(array01, array02);
        Console.WriteLine("Merged Array: " + string.Join(", ", mergedArray));
    }

    public static int[] MergeArrays(int[] array01, int[] array02)
    {
        int[] mergedArray = new int[array01.Length + array02.Length];
        int i = 0, j = 0, k = 0;

        while (i < array01.Length && j < array02.Length)
        {
            if (array01[i] < array02[j])
            {
                mergedArray[k++] = array01[i++];
            }
            else
            {
                mergedArray[k++] = array02[j++];
            }
        }

        while (i < array01.Length)
        {
            mergedArray[k++] = array01[i++];
        }

        while (j < array02.Length)
        {
            mergedArray[k++] = array02[j++];
        }

        return mergedArray;
    }
}
