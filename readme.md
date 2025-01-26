Эксперимент	batch.size	linger.ms	compression.type	buffer.memory	Source Record Write Rate (кops/sec)
1	        100	        0	        none	            33554432	    186
2	        100	        100	        snappy	            33554432	    175
3	        1000	    100	        gzip	            33554432	    182
4	        1000	    1000	    gzip	            33554432	    182
5	        10000	    10000	    gzip	            63554432	    174