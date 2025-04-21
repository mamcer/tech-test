namespace Anagrama.Cli;

public class Program
{
    public static void Main(string[] args)
    {
        Console.WriteLine($"Anagrama mario > omira: {IsAnagram("mario", "omira")}");
        Console.WriteLine($"Anagrama mario > omora: {IsAnagram("mario", "omora")}");
    }

    public static bool IsAnagram(string word1, string word2)
    {
        if (word1.Length != word2.Length)
        {
            return false;
        }


        var charCount = new int[26];

        foreach (var c in word1)
        {
            charCount[c - 'a']++;
        }

        foreach (var c in word2)
        {
            charCount[c - 'a']--;
            if (charCount[c - 'a'] < 0)
            {
                return false;
            }
        }

        return true;
    }
}