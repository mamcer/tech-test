namespace Compactor.Cli.Test;

using Compactor.Cli;

public class ProgramTest
{
    [Fact]
    public void Compactor_ReturnsCorrectResult_ForValidInput()
    {
        // Arrange
        string input = "aaaaabbbadddcc";

        // Act
        string result = Program.Compactor(input);

        // Assert
        Assert.Equal("a5b3a1d3c2", result);
    }

    [Fact]
    public void Compactor_ReturnsEmptyString_ForEmptyInput()
    {
        // Arrange
        string input = "";

        // Act
        string result = Program.Compactor(input);

        // Assert
        Assert.Equal(string.Empty, result);
    }

    [Fact]
    public void Compactor_ReturnsCorrectResult_ForSingleCharacterInput()
    {
        // Arrange
        string input = "a";

        // Act
        string result = Program.Compactor(input);

        // Assert
        Assert.Equal("a1", result);
    }

    [Fact]
    public void Compactor_ReturnsCorrectResult_ForInputWithNoRepeatingCharacters()
    {
        // Arrange
        string input = "abc";

        // Act
        string result = Program.Compactor(input);

        // Assert
        Assert.Equal("a1b1c1", result);
    }

    [Fact]
    public void Compactor_ReturnsCorrectResult_ForInputWithAllSameCharacters()
    {
        // Arrange
        string input = "aaaaa";

        // Act
        string result = Program.Compactor(input);

        // Assert
        Assert.Equal("a5", result);
    }

    [Fact]
    public void Compactor_ReturnsCorrectResult_ForMixedCaseInput()
    {
        // Arrange
        string input = "aaAA";

        // Act
        string result = Program.Compactor(input);

        // Assert
        Assert.Equal("a2A2", result);
    }

    [Fact]
    public void Compactor_ReturnsCorrectResult_ForInputWithSpecialCharacters()
    {
        // Arrange
        string input = "aa!!bb";

        // Act
        string result = Program.Compactor(input);

        // Assert
        Assert.Equal("a2!2b2", result);
    }

    [Fact]
    public void Compactor_ReturnsCorrectResult_ForInputWithNumbers()
    {
        // Arrange
        string input = "112233333";

        // Act
        string result = Program.Compactor(input);

        // Assert
        Assert.Equal("122235", result);
    }
}