namespace Merge.Cli;

internal class Program
{
    private static void Main(string[] args)
    {
        var array1 = new[] { 1, 3, 5, 7, 9 };
        var array2 = new[] { 2, 4, 6, 8, 10 };
        var mergedArray = MergeArrays(array1, array2);
        Console.WriteLine("Merged Array: " + string.Join(", ", mergedArray));
    }
}
