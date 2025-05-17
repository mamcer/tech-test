namespace Two.Cli;

public class Program
{
    public static void Main()
    {
        var target = 9;
        var numbers = new[] { 2, 7, 11, 15 };
        Console.WriteLine("The array has two numbers that sum to " + target + ": " + HasTwoSum(numbers, target));
    }

    public static bool HasTwoSum(int[] numbers, int target)
    {
        var seen = new HashSet<int>();
        foreach (var num in numbers)
        {
            int complement = target - num;
            if (seen.Contains(complement))
            {
                return true;

            }

            seen.Add(num);
        }

        return false;
    }
}