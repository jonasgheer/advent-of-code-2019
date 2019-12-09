package space;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Image {
    private int width, height;
    private ArrayList<ImageLayer> layers = new ArrayList<>();

    Image(int width, int height, int... imageData) {
        this.width = width;
        this.height = height;

        int pixelsPerLayer = width * height;

        for (int i = 0; i < imageData.length; i += pixelsPerLayer) {
            ImageLayer layer = new ImageLayer(width, height, Arrays.copyOfRange(imageData, i, i + pixelsPerLayer));
            layers.add(layer);
        }
    }

    public List<ImageLayer> getLayers() {
        return this.layers;
    }

    public int getWidth() {
        return width;
    }

    public int getHeight() {
        return height;
    }
}
