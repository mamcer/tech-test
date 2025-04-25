namespace Compactor.Cli;

using System;

public class Program
{
    public static void Main(string[] args)
    {
        var word = "aaaaabbbadddcc";
        Console.WriteLine($"compactor result for: {word} is {Compactor(word)}");
    }

    public static string Compactor(string input)
    {
        if (string.IsNullOrEmpty(input))
        {
            return string.Empty;
        }

        var current = input[0];
        var result = string.Empty;
        var count = 1;
        for (int i = 1; i < input.Length; i++)
        {
            if (input[i] == current)
            {
                count++;
                continue;
            }
            else
            {
                result += current + count.ToString();
                current = input[i];
                count = 1;
            }
        }
        result += current + count.ToString();
        return result;
    }
}