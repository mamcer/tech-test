using System;

namespace Times.Cli;

public class Program
{
    public static void Main()
    {
        char value = 'a';
        string word = "aabbccddaaeeffgg";
        Console.WriteLine($"{value} is appearing {Times(value, word)} times in the word '{word}'");
    }

    public static int Times(char value, string word)
    {
        if (word == null)
            throw new ArgumentNullException(nameof(word), "Input word cannot be null.");

        return word.Count(c => c == value);
    }
}