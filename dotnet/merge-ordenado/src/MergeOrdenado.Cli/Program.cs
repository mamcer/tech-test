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

    public static int[] MergeOrdenado(int[] array01, int[] array02)
    {
        var result = new int[array01.Length + array02.Length];
        var i = 0;
        var j = 0;
        var k = 0;

        while (i < array01.Length && j < array02.Length)
        {
            if (array01[i] < array02[j])
            {
                result[k++] = array01[i++];
            }
            else
            {
                result[k++] = array02[j++];
            }
        }

        while (i < array01.Length)
        {
            result[k++] = array01[i++];
        }

        while (j < array02.Length)
        {
            result[k++] = array02[j++];
        }

        return result;
    }
}