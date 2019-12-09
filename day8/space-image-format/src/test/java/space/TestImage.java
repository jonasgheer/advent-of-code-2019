package space;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;

import java.util.Arrays;

import static org.junit.jupiter.api.Assertions.*;

public class TestImage {

    @DisplayName("an image")
    static class TestingImage {
        int width = 3;
        int height = 2;
        int[] imageData = new int[]{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2};

        Image image = new Image(width, height, imageData);

        @Test
        @DisplayName("is instantiated with correct number of layers")
        void layers() {
            int expected = 2;
            assertEquals(expected, image.getLayers().size());
        }

        @Test
        @DisplayName("is instantiated and contains correct layer data")
        void layerData() {
            int[] expectedLayerOne = Arrays.copyOfRange(imageData, 0, 6);
            int[] expectedLayerTwo = Arrays.copyOfRange(imageData, 6, imageData.length);

            assertAll(
                    "layer",
                    () -> assertArrayEquals(expectedLayerOne, image.getLayers().get(0).getValues()),
                    () -> assertArrayEquals(expectedLayerTwo, image.getLayers().get(1).getValues())
            );
        }

        @Test
        @DisplayName("is instantiated and contains correct number of twos in layer one")
        void numberOfTwosInLayerOne() {
            ImageLayer layerOne = image.getLayers().get(0);
            int expected = layerOne.occurencesOf(2);
            assertEquals(expected, 1);
        }
    }
}
