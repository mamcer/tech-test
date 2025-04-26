namespace Maximum.Cli.Test;

using Maximum.Cli;

public class ProgramTest
{
    [Fact]
    public void MaximumDifference_ReturnsCorrectResult_ForValidInput()
    {
        // Arrange
        int[] input = { 15, 3, 6, 10 };

        // Act
        int result = Maximum.Cli.Program.MaximumDifference(input);

        // Assert
        Assert.Equal(7, result);
    }

    [Fact]
    public void MaximumDifference_ReturnsNegativeOne_ForNoValidDifference()
    {
        // Arrange
        int[] input = { 10, 9, 8, 7 };

        // Act
        int result = Maximum.Cli.Program.MaximumDifference(input);

        // Assert
        Assert.Equal(-1, result);
    }

    [Fact]
    public void MaximumDifference_ReturnsCorrectResult_ForSingleElementArray()
    {
        // Arrange
        int[] input = { 5 };

        // Act
        int result = Maximum.Cli.Program.MaximumDifference(input);

        // Assert
        Assert.Equal(-1, result);
    }

    [Fact]
    public void MaximumDifference_ReturnsCorrectResult_ForEmptyArray()
    {
        // Arrange
        int[] input = Array.Empty<int>();

        // Act
        int result = Maximum.Cli.Program.MaximumDifference(input);

        // Assert
        Assert.Equal(-1, result);
    }

    [Fact]
    public void MaximumDifference_ReturnsCorrectResult_ForArrayWithDuplicates()
    {
        // Arrange
        int[] input = { 2, 2, 2, 2 };

        // Act
        int result = Maximum.Cli.Program.MaximumDifference(input);

        // Assert
        Assert.Equal(-1, result);
    }

    [Fact]
    public void MaximumDifference_ReturnsCorrectResult_ForArrayWithIncreasingValues()
    {
        // Arrange
        int[] input = { 1, 2, 3, 4, 5 };

        // Act
        int result = Maximum.Cli.Program.MaximumDifference(input);

        // Assert
        Assert.Equal(4, result);
    }

    [Fact]
    public void MaximumDifference_ReturnsCorrectResult_ForArrayWithNegativeValues()
    {
        // Arrange
        int[] input = { -10, -20, -30, -5 };

        // Act
        int result = Maximum.Cli.Program.MaximumDifference(input);

        // Assert
        Assert.Equal(25, result);
    }
}
