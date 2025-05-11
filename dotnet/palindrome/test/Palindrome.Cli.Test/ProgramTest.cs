namespace Palindrome.Cli.Test;

using Palindrome.Cli;

public class ProgramTest
{
    [Fact]
    public void IsPalindrome_ReturnsTrue_ForPalindromeWord()
    {
        // Arrange
        string input = "racecar";

        // Act
        bool result = Program.IsPalindrome(input);

        // Assert
        Assert.True(result);
    }

    [Fact]
    public void IsPalindrome_ReturnsFalse_ForNonPalindromeWord()
    {
        // Arrange
        string input = "hello";

        // Act
        bool result = Program.IsPalindrome(input);

        // Assert
        Assert.False(result);
    }

    [Fact]
    public void IsPalindrome_ReturnsTrue_ForSingleCharacterWord()
    {
        // Arrange
        string input = "a";

        // Act
        bool result = Program.IsPalindrome(input);

        // Assert
        Assert.True(result);
    }

    [Fact]
    public void IsPalindrome_ReturnsTrue_ForEmptyString()
    {
        // Arrange
        string input = "";

        // Act
        bool result = Program.IsPalindrome(input);

        // Assert
        Assert.True(result);
    }

    [Fact]
    public void IsPalindrome_ReturnsTrue_ForPalindromeWithMixedCase()
    {
        // Arrange
        string input = "RaceCar";

        // Act
        bool result = Program.IsPalindrome(input.ToLower());

        // Assert
        Assert.True(result);
    }

    [Fact]
    public void IsPalindrome_ReturnsFalse_ForStringWithSpecialCharacters()
    {
        // Arrange
        string input = "race!car";

        // Act
        bool result = Program.IsPalindrome(input);

        // Assert
        Assert.False(result);
    }

    [Fact]
    public void IsPalindrome_ReturnsTrue_ForPalindromeWithSpaces()
    {
        // Arrange
        string input = "a man a plan a canal panama".Replace(" ", "");

        // Act
        bool result = Program.IsPalindrome(input);

        // Assert
        Assert.True(result);
    }
}