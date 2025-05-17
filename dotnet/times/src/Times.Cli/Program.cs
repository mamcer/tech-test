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
        int count = 0;
        foreach (char c in word)
        {
            if (c == value)
            {
                count++;
            }
        }
 
        return count;
    }
}