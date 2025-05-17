namespace Two.Cli.Test;

public class Program
{
    [Fact]
    public void HasTwoSum_ReturnsTrue_WhenPairExists()
    {
        // Arrange
        int[] numbers = { 2, 7, 5, 15 };
        int target = 9;

        // Act
        bool result = Two.Cli.Program.HasTwoSum(numbers, target);

        // Assert
        Assert.True(result);
    }

    [Fact]
    public void HasTwoSum_ReturnsFalse_WhenNoPairExists()
    {
        // Arrange
        int[] numbers = { 1, 2, 3, 4 };
        int target = 10;

        // Act
        bool result = Two.Cli.Program.HasTwoSum(numbers, target);

        // Assert
        Assert.False(result);
    }

    [Fact]
    public void HasTwoSum_ReturnsTrue_WhenPairIsSameElementTwice()
    {
        // Arrange
        int[] numbers = { 4, 4 };
        int target = 8;

        // Act
        bool result = Two.Cli.Program.HasTwoSum(numbers, target);

        // Assert
        Assert.True(result);
    }

    [Fact]
    public void HasTwoSum_ReturnsFalse_ForEmptyArray()
    {
        // Arrange
        int[] numbers = { };
        int target = 5;

        // Act
        bool result = Two.Cli.Program.HasTwoSum(numbers, target);

        // Assert
        Assert.False(result);
    }

    [Fact]
    public void HasTwoSum_ReturnsFalse_ForSingleElementArray()
    {
        // Arrange
        int[] numbers = { 5 };
        int target = 10;

        // Act
        bool result = Two.Cli.Program.HasTwoSum(numbers, target);

        // Assert
        Assert.False(result);
    }
}
