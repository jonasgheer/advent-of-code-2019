package space;

import java.io.IOError;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.Comparator;
import java.util.stream.Collectors;
import java.util.stream.Stream;

public class Main {
    public static void main(String[] args) throws IOException {
        String[] data = Files.readString(Paths.get("src/main/resources/image.txt")).split("");

        int[] imageData = Arrays.stream(data).mapToInt(Integer::parseInt).toArray();

        Image image = new Image(25, 6, imageData);

        int currMinZeroes = image.getLayers().get(0).occurencesOf(0);
        int index = 0;
        for (int i = 0; i < image.getLayers().size(); i++) {
            int currZeroes = image.getLayers().get(i).occurencesOf(0);
            if (currZeroes < currMinZeroes) {
                currMinZeroes = currZeroes;
                index = i;
            }
        }

        int ones = image.getLayers().get(index).occurencesOf(1);
        int twos = image.getLayers().get(index).occurencesOf(2);

        System.out.println("ones * twos = " + ones * twos);
    }
}
