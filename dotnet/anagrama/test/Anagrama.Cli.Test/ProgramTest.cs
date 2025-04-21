namespace Anagrama.Cli.Test;

public class ProgramTest
{
    [Fact]
    public void IsAnagram_ReturnsTrue_ForValidAnagrams()
    {
        // Arrange
        string word1 = "mario";
        string word2 = "omira";

        // Act
        bool result = Program.IsAnagram(word1, word2);

        // Assert
        Assert.True(result);
    }

    [Fact]
    public void IsAnagram_ReturnsFalse_ForDifferentLengthWords()
    {
        // Arrange
        string word1 = "mario";
        string word2 = "omar";

        // Act
        bool result = Program.IsAnagram(word1, word2);

        // Assert
        Assert.False(result);
    }

    [Fact]
    public void IsAnagram_ReturnsFalse_ForNonAnagrams()
    {
        // Arrange
        string word1 = "mario";
        string word2 = "omora";

        // Act
        bool result = Program.IsAnagram(word1, word2);

        // Assert
        Assert.False(result);
    }

    [Fact]
    public void IsAnagram_ReturnsTrue_ForEmptyStrings()
    {
        // Arrange
        string word1 = "";
        string word2 = "";

        // Act
        bool result = Program.IsAnagram(word1, word2);

        // Assert
        Assert.True(result);
    }
}