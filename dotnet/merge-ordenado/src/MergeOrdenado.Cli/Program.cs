namespace MergeOrdenado.Cli;

public class Program
{
    public static void Main(string[] args)
    {
        var array01 = new int[] { 1, 3, 5, 7, 9 };
        var array02 = new int[] { 2, 4, 6, 8, 10 };

        Console.WriteLine("array 01: " + string.Join(", ", array01));
        Console.WriteLine("array 02: " + string.Join(", ", array01));
        Console.WriteLine($"merge ordenado result: {MergeOrdenado(array01, array02)}");
    }
}