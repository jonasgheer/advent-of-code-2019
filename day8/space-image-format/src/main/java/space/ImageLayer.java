package space;

public class ImageLayer {
    private int width, height;
    private int[] values;

    ImageLayer(int width, int height, int... values) {
        this.width = width;
        this.height = height;
        this.values = values;
    }

    public int occurencesOf(int value) {
        int count = 0;
        for (int n : values) {
            if (n == value) {
                count++;
            }
        }
        return count;
    }

    public int getWidth() {
        return width;
    }

    public int getHeight() {
        return height;
    }

    public int[] getValues() {
        return values;
    }
}
