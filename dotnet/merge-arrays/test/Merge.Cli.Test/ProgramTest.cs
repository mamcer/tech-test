namespace Merge.Cli.Test;

using Merge.Cli;

public class ProgramTest
{
    [Fact]
    public void MergeArrays_ReturnsMergedArray_ForTwoSortedArrays()
    {
        // Arrange
        int[] array01 = { 1, 3, 5 };
        int[] array02 = { 2, 4, 6 };

        // Act
        int[] result = Program.MergeArrays(array01, array02);

        // Assert
        Assert.Equal(new[] { 1, 2, 3, 4, 5, 6 }, result);
    }

    [Fact]
    public void MergeArrays_ReturnsFirstArray_WhenSecondArrayIsEmpty()
    {
        // Arrange
        int[] array01 = { 1, 3, 5 };
        int[] array02 = Array.Empty<int>();

        // Act
        int[] result = Program.MergeArrays(array01, array02);

        // Assert
        Assert.Equal(new[] { 1, 3, 5 }, result);
    }

    [Fact]
    public void MergeArrays_ReturnsSecondArray_WhenFirstArrayIsEmpty()
    {
        // Arrange
        int[] array01 = Array.Empty<int>();
        int[] array02 = { 2, 4, 6 };

        // Act
        int[] result = Program.MergeArrays(array01, array02);

        // Assert
        Assert.Equal(new[] { 2, 4, 6 }, result);
    }

    [Fact]
    public void MergeArrays_ReturnsEmptyArray_WhenBothArraysAreEmpty()
    {
        // Arrange
        int[] array01 = Array.Empty<int>();
        int[] array02 = Array.Empty<int>();

        // Act
        int[] result = Program.MergeArrays(array01, array02);

        // Assert
        Assert.Empty(result);
    }

    [Fact]
    public void MergeArrays_ReturnsMergedArray_ForArraysWithDuplicateValues()
    {
        // Arrange
        int[] array01 = { 1, 3, 5 };
        int[] array02 = { 1, 3, 5 };

        // Act
        int[] result = Program.MergeArrays(array01, array02);

        // Assert
        Assert.Equal(new[] { 1, 1, 3, 3, 5, 5 }, result);
    }

    [Fact]
    public void MergeArrays_ReturnsMergedArray_ForArraysWithNegativeValues()
    {
        // Arrange
        int[] array01 = { -10, -5, 0 };
        int[] array02 = { -8, -3, 2 };

        // Act
        int[] result = Program.MergeArrays(array01, array02);

        // Assert
        Assert.Equal(new[] { -10, -8, -5, -3, 0, 2 }, result);
    }

    [Fact]
    public void MergeArrays_ReturnsMergedArray_ForArraysWithDifferentLengths()
    {
        // Arrange
        int[] array01 = { 1, 3 };
        int[] array02 = { 2, 4, 6, 8 };

        // Act
        int[] result = Program.MergeArrays(array01, array02);

        // Assert
        Assert.Equal(new[] { 1, 2, 3, 4, 6, 8 }, result);
    }

        [Fact]
    public void MergeArrays_ThrowsArgumentNullException_WhenFirstArrayIsNull()
    {
        // Arrange
        int[] array01 = null;
        int[] array02 = { 1, 2, 3 };

        // Act & Assert
        Assert.Throws<ArgumentNullException>(() => Program.MergeArrays(array01, array02));
    }

    [Fact]
    public void MergeArrays_ThrowsArgumentNullException_WhenSecondArrayIsNull()
    {
        // Arrange
        int[] array01 = { 1, 2, 3 };
        int[] array02 = null;

        // Act & Assert
        Assert.Throws<ArgumentNullException>(() => Program.MergeArrays(array01, array02));
    }
}