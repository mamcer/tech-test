namespace Times.Cli.Test;

public class TimesTest
{
    [Fact]
    public void Times_ReturnsCorrectCount_ForCharacterPresentMultipleTimes()
    {
        // Arrange
        char value = 'a';
        string word = "aabbccddaaeeffgg";

        // Act
        int result = Times.Cli.Program.Times(value, word);

        // Assert
        Assert.Equal(4, result);
    }

    [Fact]
    public void Times_ReturnsZero_WhenCharacterNotPresent()
    {
        // Arrange
        char value = 'z';
        string word = "aabbccddaaeeffgg";

        // Act
        int result = Times.Cli.Program.Times(value, word);

        // Assert
        Assert.Equal(0, result);
    }

    [Fact]
    public void Times_ReturnsZero_ForEmptyString()
    {
        // Arrange
        char value = 'a';
        string word = "";

        // Act
        int result = Times.Cli.Program.Times(value, word);

        // Assert
        Assert.Equal(0, result);
    }

    [Fact]
    public void Times_ThrowsArgumentNullException_ForNullString()
    {
        // Arrange
        char value = 'a';
        string word = null!;

        // Act & Assert
        Assert.Throws<ArgumentNullException>(() => Times.Cli.Program.Times(value, word!));
    }

    [Fact]
    public void Times_IsCaseSensitive()
    {
        // Arrange
        char value = 'a';
        string word = "AaAa";

        // Act
        int result = Times.Cli.Program.Times(value, word);

        // Assert
        Assert.Equal(2, result);
    }
}
